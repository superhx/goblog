var gulp = require('gulp');
var uglify = require('gulp-uglify');
var shell = require('gulp-shell');
var minifyCSS = require('gulp-minify-css');
var react = require('gulp-react');
var concat = require('gulp-concat');
var order = require("gulp-order");
var merge = require('merge-stream');
var watch = require('gulp-watch');


var DATA_JS_SRC = 'src/data/js/';
var DATA_JS_DEST = 'data/js/';
var DATA_CSS_SRC = 'src/data/css/';
var DATA_CSS_DEST = 'data/css';
var TMPL = 'template/'
var LIB_JS_DIR = 'src/lib/js/';
var BLOG_DIR = '/Users/Lancer/blog_preview/';

gulp.task('js', function() {
    var jsx = gulp.src(DATA_JS_SRC + '*.jsx')
        .pipe(react()).on('error', console.error.bind(console));
    var js = gulp.src(DATA_JS_SRC + '*.js');

    merge(jsx, js).pipe(uglify())
        .pipe(concat('goblog.min.js'))
        .pipe(gulp.dest(DATA_JS_DEST))
})

gulp.task('css', function() {
    gulp.src(DATA_CSS_SRC + '*.css')
        .pipe(minifyCSS())
        .pipe(concat('goblog.min.css'))
        .pipe(gulp.dest(DATA_CSS_DEST))
})

gulp.task('lib', function() {
    gulp.src(LIB_JS_DIR + '*.js')
        .pipe(gulp.dest(DATA_JS_DEST));
})
gulp.task('template', shell.task([
    'rm -f category.json',
    'goblog g'
], {
    'cwd': BLOG_DIR
}))

gulp.task('init', function(){
    gulp.src('').pipe(shell([
        'goblog i'
    ], {
        'cwd': BLOG_DIR
    }));
})

gulp.task('monitor', function() {
    gulp.watch(DATA_JS_SRC + '*.js', ['js'])
    gulp.watch(DATA_JS_SRC + '*.jsx', ['js'])
    gulp.watch(DATA_CSS_SRC + '*.css', ['css'])
    gulp.watch('data/*/*', ['init'])
    gulp.watch(TMPL + '*.htm', ['template'])
})

gulp.task('default', ['lib', 'js', 'css', 'template', 'init', 'monitor'])
