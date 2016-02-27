var gulp = require('gulp');
var shell = require('gulp-shell');
var rename = require('gulp-rename');
var concat = require('gulp-concat');
var watch = require('gulp-watch');
var minifyCSS = require('gulp-minify-css');
var browserify = require('gulp-browserify');
var reactify = require('reactify');

var jsSrc = 'src/data/js/';
var jsDest = 'data/js/';
gulp.task('js', function() {
  gulp
    .src(jsSrc + 'main.jsx')
    .pipe(browserify({
      transform: ['reactify'],
      extensions: ['jsx']
    }))
    .on('error', function(err){
      console.log(err.message);
      this.emit('end');
    })
    .pipe(rename('blog.min.js'))
    .pipe(gulp.dest(jsDest));

});

var cssSrc = 'src/data/css/';
var cssDest = 'data/css/';
gulp.task('css', function() {
  gulp.src(cssSrc + '*.css')
    .pipe(minifyCSS())
    .pipe(concat('blog.min.css'))
    .pipe(gulp.dest(cssDest));
});

var jsLibSrc = 'src/lib/js/';
gulp.task('lib', function() {
  var jsLibs = [
    'react-with-addons.min.js',
    'react-dom.min.js',
    'jquery.min.js',
    'scrollArea.js',
    'date.js',
    'highlight.pack.js'
  ];

  gulp.src(jsLibs.map(function(e) { return jsLibSrc + e; }))
    .pipe(concat('lib.min.js'))
    .pipe(gulp.dest(jsDest));
  gulp.src('src/lib/js/*')
    .pipe(gulp.dest(jsDest));
  gulp.src('src/lib/css/*')
    .pipe(gulp.dest(cssDest));
});

var blogDir = '/Users/Lancer/blog_preview';
gulp.task('template', shell.task([
  'rm -f catalogue.json',
  'goblog g'
], {
  'cwd': blogDir
}));

gulp.task('init', shell.task(
  ['goblog i'], {
    'cwd': blogDir
  }
));

var TMPL = 'template/';
gulp.task('monitor', function() {
  gulp.watch(jsSrc + '/**/*.jsx', ['js']);
  gulp.watch(cssSrc + '*.css', ['css']);
  gulp.watch('data/*/*', ['init']);
  gulp.watch(TMPL + '*.htm', ['template']);
});

gulp.task('default', ['lib', 'js', 'css', 'template', 'init', 'monitor']);
