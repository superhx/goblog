var gulp = require('gulp');
var shell = require('gulp-shell');
var rename = require('gulp-rename');
var concat = require('gulp-concat');
var watch = require('gulp-watch');
var webpack = require('webpack-stream');

var jsSrc = 'src/data/js/';
var jsDest = 'data/js/';
gulp.task('js', shell.task(['webpack --process']));

var cssSrc = 'src/data/css/';
var cssDest = 'data/css/';
gulp.task('css', function() {
  gulp.src(cssSrc + '*.css')
    .pipe(concat('blog.min.css'))
    .pipe(gulp.dest(cssDest));
});

var jsLibSrc = 'src/lib/js/';
gulp.task('lib', function() {
  var jsLibs = [
    'scrollArea.js',
  ];

  gulp.src(jsLibs.map(function(e) { return jsLibSrc + e; }))
    .pipe(concat('lib.min.js'))
    .pipe(gulp.dest(jsDest));
  gulp.src('src/lib/css/*')
    .pipe(concat('lib.min.css'))
    .pipe(gulp.dest(cssDest));
});

var blogDir = '/Users/robin/Blog';
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
// gulp.task('default', ['js']);
